package tagcloud

type TagCloud struct {
	storage map[string]int
}

type TagStat struct {
	Tag             string
	OccurrenceCount int
}

func New() *TagCloud {
	return &TagCloud{
		storage: make(map[string]int),
	}
}

func (c TagCloud) AddTag(tag string) {
	res, exists := c.storage[tag]
	if !exists {
		c.storage[tag] = 1
	} else {
		c.storage[tag] = res + 1
	}
}

func (c TagCloud) TopN(n int) []TagStat {
	return c.findKMax(n)
}
