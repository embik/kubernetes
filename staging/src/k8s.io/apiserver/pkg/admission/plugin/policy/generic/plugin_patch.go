package generic

import (
	"k8s.io/client-go/informers"
	coreinformers "k8s.io/client-go/informers/core/v1"
)

func (c *Plugin[H]) SetNamespaceInformer(i coreinformers.NamespaceInformer) {
	c.namespaceInformer = i
}

func (c *Plugin[H]) SetInformerFactory(f informers.SharedInformerFactory) {
	c.informerFactory = f
}

func (c *Plugin[H]) SetSourceFactory(s sourceFactory[H]) {
	c.sourceFactory = s
}
