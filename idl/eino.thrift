namespace go eino


struct CallAssistantAgentReq{
    string user_uuid
    string content
}

struct CallAssistantAgentResp{
    string reply
}

service EinoService{
    CallAssistantAgentResp CallAssistantAgent(1: CallAssistantAgentReq req)
}


