package optional

type Option[T any] struct {
	value T
	set   bool
}

func None[T any]() Option[T] {
	return Option[T]{}
}

func Some[T any](value T) Option[T] {
	return Option[T]{value: value, set: true}
}

func FromPointer[T any](value *T) Option[T] {
	if value == nil {
		return None[T]()
	}

	return Some(*value)
}

func (o Option[T]) Get() (T, bool) {
	return o.value, o.set
}

func (o Option[T]) IsSet() bool {
	return o.set
}

func (o Option[T]) IsEmpty() bool {
	return !o.set
}

func (o Option[T]) Pointer() *T {
	if !o.set {
		return nil
	}

	value := o.value
	return &value
}
