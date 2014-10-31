package characteristic

type Int struct {
    *Number
}

func NewInt(value, min, max, steps int) *Int {
    number := NewNumber(value, min, max, steps, FormatInt)
    return &Int{number}
}

func (i *Int) SetFloat(value int) {
    i.SetValue(value)
}

func (i *Int) SetMin(value int) {
    i.SetMinValue(value)
}

func (i *Int) SetMax(value int) {
    i.SetMaxValue(value)
}

func (i *Int) SetMinStep(value int) {
    i.SetMinStepValue(value)
}

func (i *Int) IntValue() int {
    return i.GetValue().(int)
}

func (i *Int) Min() int {
    return i.GetMinValue().(int)
}

func (i *Int) Max() int {
    return i.GetMaxValue().(int)
}

func (i *Int) MinStep() int {
    return i.GetMinStepValue().(int)
}