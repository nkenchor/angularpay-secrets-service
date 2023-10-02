package entity

type SecretStruct struct {
	Reference         string `bson:"reference" json:"reference"`
	CreatedOn         string `bson:"created_on" json:"created_on"`
	LastModified      string `bson:"last_modified" json:"last_modified"`
	ServiceReference string `bson:"service_reference" json:"service_reference"`
	Name              string `bson:"name" json:"name" validate:"required,min=2"`
	Value             string `bson:"value" json:"value" validate:"required"`
}