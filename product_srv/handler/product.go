package handler

import (
	"context"
	"fmt"
	product_srv "mall/product_srv/proto/product"
)

type Product struct{}

func (p *Product) ProductList(ctx context.Context, in *product_srv.ProductsRequest, out *product_srv.ProductsResponse) error {
	currentPage := in.CurrentPage
	pageSize := in.PageSize
	fmt.Println(currentPage, pageSize)
	offset := (currentPage - 1) * pageSize
	fmt.Println(offset)

	//	\\数据逻辑处理部分

	out.Code = 200
	out.Msg = "成功！"
	//out.Products = products_rep
	//out.Total = count
	out.PageSize = pageSize
	out.Current = currentPage
	return nil
}

func (p *Product) ProductAdd(ctx context.Context, in *product_srv.ProductAddRequest, out *product_srv.ProductAddResponse) error {
	name := in.Name
	price := in.Price
	num := in.Num
	unit := in.Unit
	pic := in.Pic
	desc := in.Desc
	fmt.Println(name, price, num, unit, pic, desc)
	//	\\数据逻辑处理部分

	out.Code = 200
	out.Msg = "添加商品成功"
	return nil
}
