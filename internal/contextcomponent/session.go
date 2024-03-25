package contextcomponent

import (
	"github.com/Alp4ka/classifier-aaS/internal/contextcomponent/repository"
	"github.com/Alp4ka/classifier-aaS/internal/schema"
)

type Session struct {
	Model *repository.Session
	Tree  schema.Tree
}
