"""
Normalização de dados
"""
from datetime import datetime
from typing import Any, List, Dict


class DataNormalizer:
    """Normaliza dados recebidos da API"""
    
    @staticmethod
    def normalize(raw_data: Any, metadata: Dict = None) -> List[Dict]:
        """
        Normaliza dados adicionando metadata
        
        Args:
            raw_data: Dados brutos da API
            metadata: Metadata adicional
            
        Returns:
            Lista de dados normalizados
        """
        normalized = []
        
        # Converte para lista se necessário
        items = raw_data if isinstance(raw_data, list) else [raw_data]
        
        for item in items:
            normalized_item = {
                "data": item,
                "metadata": {
                    "timestamp": datetime.utcnow().isoformat(),
                    "source": "api_collector",
                    "version": "1.0",
                    **(metadata or {})
                }
            }
            normalized.append(normalized_item)
        
        return normalized