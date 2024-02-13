import { ChatState } from '../constants/chats';

export type Chat = {
  sender: string;
  type: string;
  content: string;
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