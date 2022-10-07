from typing import Union

from fastapi import FastAPI, Request

app = FastAPI()


@app.post("/")
async def read_root(request: Request):
    data = await request.json()
    headers = request.headers
    print(headers)
    print(data)
    return data


@app.get("/items/{item_id}")
def read_item(item_id: int, q: Union[str, None] = None):
    return {"item_id": item_id, "q": q}
