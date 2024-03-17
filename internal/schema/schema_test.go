package schema

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDescription_UnmarshalJSON(t *testing.T) {
	data := `
[
  {
    "id": "77c7f51f-b22f-46c9-8db0-9961be07be8c",
    "type": "start",
    "nextID": "274c02b8-c679-48e3-8e20-a3a750556e26",
    "nextErrorID": null,
    "data": {
      "zhopaInt": 123,
      "zhopaString": "456"
    }
  },
  {
    "id": "274c02b8-c679-48e3-8e20-a3a750556e26",
    "type": "finish",
    "nextID": null,
    "nextErrorID": null,
    "data": {
      "zhopaInt": 1337
    }
  }
]
`
	var desc Description
	err := json.Unmarshal([]byte(data), &desc)
	if err != nil {
		t.Fatal(err)
	}

	{
		val, err := json.Marshal(desc)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(string(val))
	}

	require.Len(t, desc.nodes, 2)
	require.Equal(t, desc.nodes[0].ID.String(), "77c7f51f-b22f-46c9-8db0-9961be07be8c")
	require.Equal(t, desc.nodes[0].Type, NodeTypeStart)
	require.Equal(t, desc.nodes[0].NextID.UUID.String(), "274c02b8-c679-48e3-8e20-a3a750556e26")
	require.True(t, desc.nodes[0].NextID.Valid)
	require.False(t, desc.nodes[0].NextErrorID.Valid)
	require.Len(t, desc.nodes[0].Data, 2)

	require.Equal(t, desc.nodes[1].ID.String(), "274c02b8-c679-48e3-8e20-a3a750556e26")
	require.Equal(t, desc.nodes[1].Type, NodeTypeFinish)
	require.Len(t, desc.nodes[1].Data, 1)
	require.False(t, desc.nodes[1].NextID.Valid)
	require.False(t, desc.nodes[1].NextErrorID.Valid)
}

func TestDescription_MapAndValidate(t *testing.T) {
	t.Run("should map and validate", func(t *testing.T) {
		data := `
[
  {
    "id": "77c7f51f-b22f-46c9-8db0-9961be07be8c",
    "type": "start",
    "nextID": "274c02b8-c679-48e3-8e20-a3a750556e26",
    "nextErrorID": null,
    "data": {
      "zhopaInt": 123,
      "zhopaString": "456"
    }
  },
  {
    "id": "274c02b8-c679-48e3-8e20-a3a750556e26",
    "type": "finish",
    "nextID": null,
    "nextErrorID": null,
    "data": {
      "zhopaInt": 1337
    }
  }
]
`
		var desc Description
		err := json.Unmarshal([]byte(data), &desc)
		if err != nil {
			t.Fatal(err)
		}

		mapping, err := desc.MapAndValidate()
		if err != nil {
			t.Fatal(err)
		}

		require.Len(t, mapping, 2)
	})

	t.Run("unrelated next link", func(t *testing.T) {
		data := `
[
  {
    "id": "77c7f51f-b22f-46c9-8db0-9961be07be8c",
    "type": "start",
    "nextID": "8539add9-aa0c-4ab5-9200-00affdd64bc3",
    "nextErrorID": null,
    "data": {
      "zhopaInt": 123,
      "zhopaString": "456"
    }
  },
  {
    "id": "274c02b8-c679-48e3-8e20-a3a750556e26",
    "type": "finish",
    "nextID": null,
    "nextErrorID": null,
    "data": {
      "zhopaInt": 1337
    }
  }
]
`
		var desc Description
		err := json.Unmarshal([]byte(data), &desc)
		if err != nil {
			t.Fatal(err)
		}

		_, err = desc.MapAndValidate()
		require.Error(t, err)
	})

	t.Run("unrelated error next link", func(t *testing.T) {
		data := `
[
  {
    "id": "77c7f51f-b22f-46c9-8db0-9961be07be8c",
    "type": "start",
    "nextID": null,
    "nextErrorID": "8539add9-aa0c-4ab5-9200-00affdd64bc3",
    "data": {
      "zhopaInt": 123,
      "zhopaString": "456"
    }
  },
  {
    "id": "274c02b8-c679-48e3-8e20-a3a750556e26",
    "type": "finish",
    "nextID": null,
    "nextErrorID": null,
    "data": {
      "zhopaInt": 1337
    }
  }
]
`
		var desc Description
		err := json.Unmarshal([]byte(data), &desc)
		if err != nil {
			t.Fatal(err)
		}

		_, err = desc.MapAndValidate()
		require.Error(t, err)
	})

	t.Run("no start node", func(t *testing.T) {
		data := `
[
  {
    "id": "274c02b8-c679-48e3-8e20-a3a750556e26",
    "type": "finish",
    "nextID": null,
    "nextErrorID": null,
    "data": {
      "zhopaInt": 1337
    }
  }
]
`
		var desc Description
		err := json.Unmarshal([]byte(data), &desc)
		if err != nil {
			t.Fatal(err)
		}

		_, err = desc.MapAndValidate()
		require.Error(t, err)
	})

	t.Run("no finish node", func(t *testing.T) {
		data := `
[
  {
    "id": "274c02b8-c679-48e3-8e20-a3a750556e26",
    "type": "start",
    "nextID": null,
    "nextErrorID": null,
    "data": {
      "zhopaInt": 1337
    }
  }
]
`
		var desc Description
		err := json.Unmarshal([]byte(data), &desc)
		if err != nil {
			t.Fatal(err)
		}

		_, err = desc.MapAndValidate()
		require.Error(t, err)
	})

	t.Run("finish has next", func(t *testing.T) {
		data := `
[
  {
    "id": "77c7f51f-b22f-46c9-8db0-9961be07be8c",
    "type": "start",
    "nextID": "274c02b8-c679-48e3-8e20-a3a750556e26",
    "nextErrorID": null,
    "data": {
      "zhopaInt": 123,
      "zhopaString": "456"
    }
  },
  {
    "id": "274c02b8-c679-48e3-8e20-a3a750556e26",
    "type": "finish",
    "nextID": "8539add9-aa0c-4ab5-9200-00affdd64bc3",
    "nextErrorID": null,
    "data": {
      "zhopaInt": 1337
    }
  }
]
`
		var desc Description
		err := json.Unmarshal([]byte(data), &desc)
		if err != nil {
			t.Fatal(err)
		}

		_, err = desc.MapAndValidate()
		require.Error(t, err)
	})

	t.Run("finish has fail path", func(t *testing.T) {
		data := `
[
  {
    "id": "77c7f51f-b22f-46c9-8db0-9961be07be8c",
    "type": "start",
    "nextID": "274c02b8-c679-48e3-8e20-a3a750556e26",
    "nextErrorID": null,
    "data": {
      "zhopaInt": 123,
      "zhopaString": "456"
    }
  },
  {
    "id": "274c02b8-c679-48e3-8e20-a3a750556e26",
    "type": "finish",
    "nextID": null,
    "nextErrorID": "8539add9-aa0c-4ab5-9200-00affdd64bc3",
    "data": {
      "zhopaInt": 1337
    }
  }
]
`
		var desc Description
		err := json.Unmarshal([]byte(data), &desc)
		if err != nil {
			t.Fatal(err)
		}

		_, err = desc.MapAndValidate()
		require.Error(t, err)
	})

	t.Run("duplicates", func(t *testing.T) {
		data := `
[
  {
    "id": "77c7f51f-b22f-46c9-8db0-9961be07be8c",
    "type": "start",
    "nextID": "274c02b8-c679-48e3-8e20-a3a750556e26",
    "nextErrorID": null,
    "data": {
      "zhopaInt": 123,
      "zhopaString": "456"
    }
  },
  {
    "id": "77c7f51f-b22f-46c9-8db0-9961be07be8c",
    "type": "start",
    "nextID": "274c02b8-c679-48e3-8e20-a3a750556e26",
    "nextErrorID": null,
    "data": {
      "zhopaInt": 123,
      "zhopaString": "456"
    }
  },
  {
    "id": "274c02b8-c679-48e3-8e20-a3a750556e26",
    "type": "finish",
    "nextID": "8539add9-aa0c-4ab5-9200-00affdd64bc3",
    "nextErrorID": null,
    "data": {
      "zhopaInt": 1337
    }
  }
]
`
		var desc Description
		err := json.Unmarshal([]byte(data), &desc)
		if err != nil {
			t.Fatal(err)
		}

		_, err = desc.MapAndValidate()
		require.Error(t, err)
	})

	t.Run("much start nodes", func(t *testing.T) {
		data := `
[
  {
    "id": "77c7f51f-b22f-46c9-8db0-9961be07be8c",
    "type": "start",
    "nextID": "274c02b8-c679-48e3-8e20-a3a750556e26",
    "nextErrorID": null,
    "data": {
      "zhopaInt": 123,
      "zhopaString": "456"
    }
  },
  {
    "id": "6ec27a64-4cb5-41d2-82ec-e058f78d6be2",
    "type": "start",
    "nextID": "274c02b8-c679-48e3-8e20-a3a750556e26",
    "nextErrorID": null,
    "data": {
      "zhopaInt": 123,
      "zhopaString": "456"
    }
  },
  {
    "id": "274c02b8-c679-48e3-8e20-a3a750556e26",
    "type": "finish",
    "nextID": "8539add9-aa0c-4ab5-9200-00affdd64bc3",
    "nextErrorID": null,
    "data": {
      "zhopaInt": 1337
    }
  }
]
`
		var desc Description
		err := json.Unmarshal([]byte(data), &desc)
		if err != nil {
			t.Fatal(err)
		}

		_, err = desc.MapAndValidate()
		require.Error(t, err)
	})

	t.Run("much finish nodes", func(t *testing.T) {
		data := `
[
  {
    "id": "77c7f51f-b22f-46c9-8db0-9961be07be8c",
    "type": "start",
    "nextID": "274c02b8-c679-48e3-8e20-a3a750556e26",
    "nextErrorID": null,
    "data": {
      "zhopaInt": 123,
      "zhopaString": "456"
    }
  },
  {
    "id": "6ec27a64-4cb5-41d2-82ec-e058f78d6be2",
    "type": "finish",
    "nextID": "274c02b8-c679-48e3-8e20-a3a750556e26",
    "nextErrorID": null,
    "data": {
      "zhopaInt": 123,
      "zhopaString": "456"
    }
  },
  {
    "id": "274c02b8-c679-48e3-8e20-a3a750556e26",
    "type": "finish",
    "nextID": "8539add9-aa0c-4ab5-9200-00affdd64bc3",
    "nextErrorID": null,
    "data": {
      "zhopaInt": 1337
    }
  }
]
`
		var desc Description
		err := json.Unmarshal([]byte(data), &desc)
		if err != nil {
			t.Fatal(err)
		}

		_, err = desc.MapAndValidate()
		require.Error(t, err)
	})
}
