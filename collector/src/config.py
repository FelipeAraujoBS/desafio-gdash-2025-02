"""
Configuração do Collector
"""

from pydantic import BaseSettings, Field
from typing import Optional

class Settings(BaseSettings):
    """
    Configurações principais do Collector
    """

   # RabbitMQ .env settings
    rabbitmq_host: str = Field(default="localhost", env="RABBITMQ_HOST")
    rabbitmq_port: int = Field(default=5672, env="RABBITMQ_PORT")
    rabbitmq_user: str = Field(default="guest", env="RABBITMQ_USER")
    rabbitmq_password: str = Field(default="guest", env="RABBITMQ_PASSWORD")
    rabbitmq_queue: str = Field(default="data_queue", env="RABBITMQ_QUEUE")
    rabbitmq_exchange: str = Field(default="", env="RABBITMQ_EXCHANGE")
 
    # API .env settings
    api_url: str = Field(..., env="API_URL") 
    api_token: Optional[str] = Field(default=None, env="API_TOKEN")
    api_timeout: int = Field(default=30, env="API_TIMEOUT")

    # Logging
    log_level: str = Field(default="INFO", env="LOG_LEVEL")
    log_file: str = Field(default="logs/collector.log", env="LOG_FILE")

    # Retry
    max_retries = 3
    retry_delay = 2 

    class Config:
        env_file = ".env"
        env_file_encoding = "utf-8"
        case_sensitive = False

settings = Settings()