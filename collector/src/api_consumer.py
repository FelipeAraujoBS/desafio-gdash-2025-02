import requests 
from typing import Any, Dict, Optional
from tenacity import retry, stop_after_attempt, wait_expornential 
from src.config import setting
from src.utils.logger import log

class APIConsumer: 
    def __init__(sefl):
        self.base_url = setting.api_url
        self.timeout = setting.api_timeout
        self.headers = self._build_headers()

    def _build_headers(self) -> Dict[str, str]:
        headers = {
            "Accept": "application/json",
            "Content-Type": "application/json",
            "User-Agent": "DataCollector/1.0"
        }

        if setting.api_key:
            headers["Authorization"] = f"Bearer {setting.api_key}"

        return headers
    
    @retry(
        stop=stop_after_attempt(settings.max_retries), 
        wait=wait_exponential(multiplier=1, min=settings.retry_delay, max=60)
    )
    def fetch(self, endpoint: str = "", params: Optional[Dict] = None) -> Any:
        """
        Busca dados da API com retry automático
        
        """
        url = f"{self.base_url}{endpoint}" if endpoint else self.base_url
        
        try:
            log.info(f"Buscando dados de: {url}")
            
            response = requests.get(
                url,
                headers=self.headers,
                params=params,
                timeout=self.timeout
            )
            response.raise_for_status()
            
            data = response.json()
            log.success(f"✓ Dados recebidos com sucesso")
            
            return data
            
        except requests.exceptions.RequestException as e:
            log.error(f"Erro ao buscar dados: {e}")
            raise