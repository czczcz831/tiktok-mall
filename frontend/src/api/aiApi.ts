import api from './index';
import { AIChatReq, AIChatResp, ApiResponse } from '../types/api';

// AI聊天接口
export const aiChat = async (data: AIChatReq): Promise<ApiResponse<AIChatResp>> => {
  return api.post('/eino/chat', data);
}; 