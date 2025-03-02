import React, { useState, useEffect } from 'react';
import { 
  Container, Typography, Box, Button, TextField, Table, 
  TableBody, TableCell, TableContainer, TableHead, TableRow, 
  Paper, Modal, IconButton, Grid
} from '@mui/material';
import { Delete, Edit } from '@mui/icons-material';
import { 
  Product, 
  CreateProductReq, 
  UpdateProductReq, 
  DeleteProductReq 
} from '../types/api';
import { 
  getProductList, 
  createProduct, 
  updateProduct, 
  deleteProduct 
} from '../api/productApi';
import { toast } from 'react-toastify';

// 模态框样式
const modalStyle = {
  position: 'absolute' as 'absolute',
  top: '50%',
  left: '50%',
  transform: 'translate(-50%, -50%)',
  width: 400,
  bgcolor: 'background.paper',
  border: '2px solid #000',
  boxShadow: 24,
  p: 4,
};

const AdminPage: React.FC = () => {
  // 状态管理
  const [products, setProducts] = useState<Product[]>([]);
  const [openCreateModal, setOpenCreateModal] = useState(false);
  const [openEditModal, setOpenEditModal] = useState(false);
  const [currentProduct, setCurrentProduct] = useState<Product | null>(null);
  const [formData, setFormData] = useState({
    name: '',
    description: '',
    price: 0,
    stock: 0
  });
  const [loading, setLoading] = useState(false);
  const [page, setPage] = useState(1);
  const [limit] = useState(10);
  const [totalProducts, setTotalProducts] = useState(0);

  // 加载产品列表
  const loadProducts = async () => {
    try {
      setLoading(true);
      const response = await getProductList({ page, limit });
      if (response.data.products) {
        setProducts(response.data.products);
        setTotalProducts(response.data.total);
      }
    } catch (error) {
      toast.error('加载产品列表失败');
      console.error('加载产品列表失败:', error);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    loadProducts();
  }, [page, limit]);

  // 处理表单输入变化
  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    
    let parsedValue: string | number = value;
    if (name === 'price' || name === 'stock') {
      parsedValue = value === '' ? 0 : parseFloat(value);
    }
    
    setFormData({
      ...formData,
      [name]: parsedValue
    });
  };

  // 打开创建产品模态框
  const handleOpenCreateModal = () => {
    setFormData({
      name: '',
      description: '',
      price: 0,
      stock: 0
    });
    setOpenCreateModal(true);
  };

  // 打开编辑产品模态框
  const handleOpenEditModal = (product: Product) => {
    setCurrentProduct(product);
    setFormData({
      name: product.name,
      description: product.description,
      price: product.price / 100, // 将分转换为元
      stock: product.stock
    });
    setOpenEditModal(true);
  };

  // 关闭模态框
  const handleCloseModals = () => {
    setOpenCreateModal(false);
    setOpenEditModal(false);
    setCurrentProduct(null);
  };

  // 创建产品
  const handleCreateProduct = async () => {
    try {
      if (!formData.name || formData.price <= 0) {
        toast.error('请填写完整的产品信息');
        return;
      }

      setLoading(true);
      const createData: CreateProductReq = {
        name: formData.name,
        description: formData.description,
        price: formData.price,
        stock: formData.stock
      };

      const response = await createProduct(createData);
      if (response.data.product) {
        toast.success('创建产品成功');
        handleCloseModals();
        loadProducts();
      }
    } catch (error) {
      toast.error('创建产品失败');
      console.error('创建产品失败:', error);
    } finally {
      setLoading(false);
    }
  };

  // 更新产品
  const handleUpdateProduct = async () => {
    try {
      if (!currentProduct || !formData.name || formData.price <= 0) {
        toast.error('请填写完整的产品信息');
        return;
      }

      setLoading(true);
      const updateData: UpdateProductReq = {
        product: {
          uuid: currentProduct.uuid,
          name: formData.name,
          description: formData.description,
          price: formData.price,
          stock: formData.stock
        }
      };

      const response = await updateProduct(updateData);
      if (response.data.product) {
        toast.success('更新产品成功');
        handleCloseModals();
        loadProducts();
      }
    } catch (error) {
      toast.error('更新产品失败');
      console.error('更新产品失败:', error);
    } finally {
      setLoading(false);
    }
  };

  // 删除产品
  const handleDeleteProduct = async (productUuid: string) => {
    if (!window.confirm('确定要删除此产品吗？')) {
      return;
    }

    try {
      setLoading(true);
      const deleteData: DeleteProductReq = {
        uuid: productUuid
      };

      const response = await deleteProduct(deleteData);
      if (response.data.uuid) {
        toast.success('删除产品成功');
        loadProducts();
      }
    } catch (error) {
      toast.error('删除产品失败');
      console.error('删除产品失败:', error);
    } finally {
      setLoading(false);
    }
  };

  return (
    <Container maxWidth="lg">
      <Box sx={{ my: 4 }}>
        <Typography variant="h4" component="h1" gutterBottom>
          产品管理
        </Typography>
        
        <Button 
          variant="contained" 
          color="primary" 
          onClick={handleOpenCreateModal}
          sx={{ mb: 3 }}
        >
          添加新产品
        </Button>

        <TableContainer component={Paper}>
          <Table>
            <TableHead>
              <TableRow>
                <TableCell>名称</TableCell>
                <TableCell>描述</TableCell>
                <TableCell align="right">价格 (¥)</TableCell>
                <TableCell align="right">库存</TableCell>
                <TableCell align="center">操作</TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {products.map((product) => (
                <TableRow key={product.uuid}>
                  <TableCell>{product.name}</TableCell>
                  <TableCell>{product.description}</TableCell>
                  <TableCell align="right">{(product.price / 100).toFixed(2)}</TableCell>
                  <TableCell align="right">{product.stock}</TableCell>
                  <TableCell align="center">
                    <IconButton 
                      color="primary" 
                      onClick={() => handleOpenEditModal(product)}
                    >
                      <Edit />
                    </IconButton>
                    <IconButton 
                      color="error" 
                      onClick={() => handleDeleteProduct(product.uuid)}
                    >
                      <Delete />
                    </IconButton>
                  </TableCell>
                </TableRow>
              ))}
              {products.length === 0 && (
                <TableRow>
                  <TableCell colSpan={5} align="center">
                    {loading ? '加载中...' : '暂无产品数据'}
                  </TableCell>
                </TableRow>
              )}
            </TableBody>
          </Table>
        </TableContainer>

        {/* 创建产品模态框 */}
        <Modal
          open={openCreateModal}
          onClose={handleCloseModals}
        >
          <Box sx={modalStyle}>
            <Typography variant="h6" component="h2" sx={{ mb: 2 }}>
              添加新产品
            </Typography>
            <Grid container spacing={2}>
              <Grid item xs={12}>
                <TextField
                  fullWidth
                  label="产品名称"
                  name="name"
                  value={formData.name}
                  onChange={handleInputChange}
                  required
                />
              </Grid>
              <Grid item xs={12}>
                <TextField
                  fullWidth
                  label="产品描述"
                  name="description"
                  value={formData.description}
                  onChange={handleInputChange}
                  multiline
                  rows={3}
                />
              </Grid>
              <Grid item xs={6}>
                <TextField
                  fullWidth
                  label="价格 (¥)"
                  name="price"
                  type="number"
                  value={formData.price}
                  onChange={handleInputChange}
                  inputProps={{ min: 0, step: 0.01 }}
                  required
                />
              </Grid>
              <Grid item xs={6}>
                <TextField
                  fullWidth
                  label="库存数量"
                  name="stock"
                  type="number"
                  value={formData.stock}
                  onChange={handleInputChange}
                  inputProps={{ min: 0 }}
                  required
                />
              </Grid>
              <Grid item xs={12}>
                <Box sx={{ display: 'flex', justifyContent: 'flex-end', mt: 2 }}>
                  <Button onClick={handleCloseModals} sx={{ mr: 1 }}>
                    取消
                  </Button>
                  <Button 
                    variant="contained" 
                    onClick={handleCreateProduct}
                    disabled={loading}
                  >
                    {loading ? '处理中...' : '创建产品'}
                  </Button>
                </Box>
              </Grid>
            </Grid>
          </Box>
        </Modal>

        {/* 编辑产品模态框 */}
        <Modal
          open={openEditModal}
          onClose={handleCloseModals}
        >
          <Box sx={modalStyle}>
            <Typography variant="h6" component="h2" sx={{ mb: 2 }}>
              编辑产品
            </Typography>
            <Grid container spacing={2}>
              <Grid item xs={12}>
                <TextField
                  fullWidth
                  label="产品名称"
                  name="name"
                  value={formData.name}
                  onChange={handleInputChange}
                  required
                />
              </Grid>
              <Grid item xs={12}>
                <TextField
                  fullWidth
                  label="产品描述"
                  name="description"
                  value={formData.description}
                  onChange={handleInputChange}
                  multiline
                  rows={3}
                />
              </Grid>
              <Grid item xs={6}>
                <TextField
                  fullWidth
                  label="价格 (¥)"
                  name="price"
                  type="number"
                  value={formData.price}
                  onChange={handleInputChange}
                  inputProps={{ min: 0, step: 0.01 }}
                  required
                />
              </Grid>
              <Grid item xs={6}>
                <TextField
                  fullWidth
                  label="库存数量"
                  name="stock"
                  type="number"
                  value={formData.stock}
                  onChange={handleInputChange}
                  inputProps={{ min: 0 }}
                  required
                />
              </Grid>
              <Grid item xs={12}>
                <Box sx={{ display: 'flex', justifyContent: 'flex-end', mt: 2 }}>
                  <Button onClick={handleCloseModals} sx={{ mr: 1 }}>
                    取消
                  </Button>
                  <Button 
                    variant="contained" 
                    onClick={handleUpdateProduct}
                    disabled={loading}
                  >
                    {loading ? '处理中...' : '更新产品'}
                  </Button>
                </Box>
              </Grid>
            </Grid>
          </Box>
        </Modal>
      </Box>
    </Container>
  );
};

export default AdminPage; 