import logging
import flask
import uuid


logger = logging.getLogger('default')


class ResponseData():
    def __init__(self, data={}, total=None, error=None, msg=None):
        self.data = data
        self.error = error
        self.msg = msg
        self.total = total

    def to_json(self):
        messages = None
        if self.error and self.msg:
            messages = []
            if type(self.msg) is dict:
                for value in self.msg.values():
                    if type(value) is list:
                        messages += value
                    else:
                        messages.append(value)
            elif type(self.msg) is str:
                messages.append(self.msg)
            else:
                messages += self.msg
        return {
            'data': self.data,
            'total': self.total,
            'error_code': self.error,
            'error_message': messages
        }
