FROM python:3.10.4

WORKDIR /app

COPY requirements.txt ./

RUN pip install --no-cache-dir -r requirements.txt

COPY . .

ENV GMP_API_KEY=<insert key here>

CMD [ "python", "./server.py" ]