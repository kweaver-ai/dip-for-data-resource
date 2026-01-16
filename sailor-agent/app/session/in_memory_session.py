from langchain_core.chat_history import BaseChatMessageHistory
from langchain_community.chat_message_histories import ChatMessageHistory
from app.session import BaseChatHistorySession


class InMemoryChatSession(BaseChatHistorySession):

    def __init__(self):
        self.message_history_session = {}

    def get_chat_history(
            self, session_id: str,
    ) -> BaseChatMessageHistory:
        if session_id in self.message_history_session:
            return self.message_history_session[session_id]
        else:
            self._add_chat_history(session_id, ChatMessageHistory())
            return self.message_history_session[session_id]

    def _add_chat_history(self, session_id: str, chat_history: BaseChatMessageHistory):
        self.message_history_session[session_id] = chat_history

    def delete_chat_history(self, session_id: str):
        if not self.message_history_session.pop(session_id, None):
            raise "%s not found in message_history_session" % session_id

    def clean_session(self):
        self.message_history_session = {}
