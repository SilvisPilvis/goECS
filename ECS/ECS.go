package ecs

// Component interface
type Component interface{}

// Entity struct
type Entity struct {
	id         int
	components map[string]Component
}

// World struct
type World struct {
	entities map[int]*Entity
	nextID   int
}

// NewWorld creates a new World
func NewWorld() *World {
	return &World{
		entities: make(map[int]*Entity),
		nextID:   0,
	}
}

// AddEntity creates a new Entity
func (w *World) AddEntity() *Entity {
	e := &Entity{
		id:         w.nextID,
		components: make(map[string]Component),
	}
	w.entities[e.id] = e
	w.nextID++
	return e
}

// AddComponent adds a component to an Entity
func (e *Entity) AddComponent(name string, component Component) {
	e.components[name] = component
}

// GetComponent retrieves a component from an Entity
func (e *Entity) GetComponent(name string) (Component, bool) {
	component, ok := e.components[name]
	return component, ok
}

// Example components
type PositionComponent struct {
	X, Y float64
}

type VelocityComponent struct {
	VX, VY float64
}

// System interface
type System interface {
	Update(world *World)
}

// MovementSystem example
type MovementSystem struct{}

func (s *MovementSystem) Update(world *World) {
	for _, entity := range world.entities {
		pos, hasPos := entity.GetComponent("position")
		vel, hasVel := entity.GetComponent("velocity")

		if hasPos && hasVel {
			p := pos.(*PositionComponent)
			v := vel.(*VelocityComponent)
			p.X += v.VX
			p.Y += v.VY
		}
	}
}
