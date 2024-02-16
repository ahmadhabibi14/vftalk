import { ChatState } from '../constants/chats';

interface Chat {
  sender: string;
  type: string;
  content: any | string | string [];
  datetime: Date;
}

export type ChatIn = {
  type: string;
  content: string;
}

export type ChatSendState = {
  state: ChatState;
  index: number;
}