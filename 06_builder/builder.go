package builder

/*
	Builder 建造者模式
*/

//Builder 生成器接口
type Builder interface {
	NewProduct()
	BuildPart1()
	BuildPart2()
	BuildPart3()
	GetResult() interface{}
}

//Director 指挥者
type Director struct {
	builder Builder
}

func NewDirector(builder Builder) *Director {
	return &Director{
		builder: builder,
	}
}

//ProductBuilder 产品建造者
type ProductBuilder struct {
	result *result
}

type result struct {
	part1 string
	part2 string
	part3 string
}

func (b *ProductBuilder) NewProduct() {
	b.result = new(result)
}

func (b *ProductBuilder) BuildPart1() {
	b.result.part1 = "part1"
}

func (b *ProductBuilder) BuildPart2() {
	b.result.part2 = "part2"
}

func (b *ProductBuilder) BuildPart3() {
	b.result.part3 = "part3"
}

func (b *ProductBuilder) GetResult() interface{} {
	return b.result
}

//Construct 构建
func (d *Director) Construct() {
	d.builder.NewProduct()
	d.builder.BuildPart1()
	d.builder.BuildPart2()
	d.builder.BuildPart3()
}


