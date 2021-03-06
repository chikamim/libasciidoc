package types

import (
	"fmt"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// ElementReferences the element references in the document
type ElementReferences map[string]interface{}

// ElementReferencesCollector the visitor that traverses the whole document structure in search for elements with an ID
type ElementReferencesCollector struct {
	ElementReferences ElementReferences
}

// NewElementReferencesCollector initializes a new ElementReferencesCollector
func NewElementReferencesCollector() *ElementReferencesCollector {
	return &ElementReferencesCollector{
		ElementReferences: ElementReferences{},
	}
}

// Visit Implements Visitable#Visit()
func (c *ElementReferencesCollector) Visit(element Visitable) error {
	switch e := element.(type) {
	case Section:
		elementID := e.Title.Attributes[AttrID]
		if elementID, ok := elementID.(string); ok {
			for i := 1; ; i++ {
				var key string
				if i == 1 {
					key = elementID
				} else {
					key = fmt.Sprintf("%s_%d", elementID, i)
				}
				if _, found := c.ElementReferences[key]; !found {
					log.Debugf("Adding element reference: %v", key)
					c.ElementReferences[key] = e.Title
					// override the element id
					e.Title.Attributes[AttrID] = key
					break
				}

			}
		} else {
			return errors.Errorf("unexpected type of element id: %T", elementID)
		}
	}
	return nil
}
