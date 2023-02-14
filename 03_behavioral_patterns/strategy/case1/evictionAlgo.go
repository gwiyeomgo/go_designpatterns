package main

type EvictionAlgo interface {
	evict(c *Cache) //퇴거시키다
}
