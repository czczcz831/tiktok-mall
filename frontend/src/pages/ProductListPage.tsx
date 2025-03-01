import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { getProductList } from '../api/productApi';
import { Product, GetProductListReq } from '../types/api';
import { useCart } from '../contexts/CartContext';

const ProductListPage: React.FC = () => {
  const [products, setProducts] = useState<Product[]>([]);
  const [total, setTotal] = useState<number>(0);
  const [loading, setLoading] = useState<boolean>(true);
  const [page, setPage] = useState<number>(1);
  const [limit] = useState<number>(10);
  const [searchName, setSearchName] = useState<string>('');
  const [minPrice, setMinPrice] = useState<string>('');
  const [maxPrice, setMaxPrice] = useState<string>('');
  const navigate = useNavigate();
  const { addToCart } = useCart();

  useEffect(() => {
    fetchProducts();
  }, [page]);

  const fetchProducts = async () => {
    setLoading(true);
    try {
      const req: GetProductListReq = {
        page,
        limit,
        name: searchName || undefined,
        min_price: minPrice ? parseFloat(minPrice) : undefined,
        max_price: maxPrice ? parseFloat(maxPrice) : undefined
      };
      const response = await getProductList(req);
      setProducts(response.data.products);
      setTotal(response.data.total);
    } catch (error) {
      console.error('Failed to fetch products:', error);
    } finally {
      setLoading(false);
    }
  };

  const handleSearch = () => {
    setPage(1); // 重置页码
    fetchProducts();
  };

  const handleAddToCart = async (productUuid: string) => {
    try {
      await addToCart(productUuid, 1);
      alert('商品已添加到购物车');
    } catch (error) {
      console.error('Failed to add to cart:', error);
      alert('添加到购物车失败');
    }
  };

  const handlePageChange = (newPage: number) => {
    setPage(newPage);
  };

  return (
    <div className="product-list-page">
      <h2>商品列表</h2>
      
      <div className="search-filters">
        <div className="search-input">
          <input
            type="text"
            placeholder="搜索商品名称"
            value={searchName}
            onChange={(e) => setSearchName(e.target.value)}
          />
        </div>
        <div className="price-filters">
          <input
            type="number"
            placeholder="最低价格"
            value={minPrice}
            onChange={(e) => setMinPrice(e.target.value)}
          />
          <span>-</span>
          <input
            type="number"
            placeholder="最高价格"
            value={maxPrice}
            onChange={(e) => setMaxPrice(e.target.value)}
          />
        </div>
        <button onClick={handleSearch}>搜索</button>
      </div>

      {loading ? (
        <div className="loading">加载中...</div>
      ) : (
        <>
          <div className="products-grid">
            {products.length > 0 ? (
              products.map((product) => (
                <div key={product.uuid} className="product-card">
                  <h3>{product.name}</h3>
                  <p className="description">{product.description}</p>
                  <p className="price">¥{product.price / 100}</p>
                  <p className="stock">库存: {product.stock}</p>
                  <div className="product-actions">
                    <button
                      onClick={() => navigate(`/product/${product.uuid}`)}
                      className="view-details"
                    >
                      查看详情
                    </button>
                    <button
                      onClick={() => handleAddToCart(product.uuid)}
                      className="add-to-cart"
                      disabled={product.stock <= 0}
                    >
                      {product.stock <= 0 ? '缺货' : '加入购物车'}
                    </button>
                  </div>
                </div>
              ))
            ) : (
              <div className="no-products">没有找到商品</div>
            )}
          </div>

          <div className="pagination">
            <button
              onClick={() => handlePageChange(page - 1)}
              disabled={page === 1}
            >
              上一页
            </button>
            <span>
              第 {page} 页，共 {Math.ceil(total / limit)} 页
            </span>
            <button
              onClick={() => handlePageChange(page + 1)}
              disabled={page >= Math.ceil(total / limit)}
            >
              下一页
            </button>
          </div>
        </>
      )}
    </div>
  );
};

export default ProductListPage; 