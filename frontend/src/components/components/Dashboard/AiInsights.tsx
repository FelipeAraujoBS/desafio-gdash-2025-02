import React from "react";
import type { Insight } from "../../../types";

const AIInsights: React.FC = () => {
  const insights: Insight[] = [
    {
      type: "Tendência",
      message:
        "A temperatura deve aumentar gradualmente ao longo do dia, atingindo o pico às 15h.",
      color: "border-blue-500",
    },
    {
      type: "Alerta",
      message:
        "Umidade baixa esperada entre 12h-16h. Recomenda-se hidratação constante.",
      color: "border-orange-500",
    },
    {
      type: "Previsão",
      message:
        "Condições ideais para atividades ao ar livre no período da manhã.",
      color: "border-green-500",
    },
  ];

  return (
    <div className="bg-white rounded-lg shadow p-6">
      <h2 className="text-xl font-semibold text-gray-800 mb-4">
        Insights de IA
      </h2>
      <div className="space-y-4">
        {insights.map((insight, index) => (
          <div key={index} className={`border-l-4 ${insight.color} pl-4 py-2`}>
            <p className="text-sm text-gray-700">
              <strong>{insight.type}:</strong> {insight.message}
            </p>
          </div>
        ))}

        <div className="mt-6 p-4 bg-blue-50 rounded-lg">
          <p className="text-xs text-gray-600 italic">
            Insights gerados por IA baseados nos dados meteorológicos atuais e
            padrões históricos.
          </p>
        </div>
      </div>
    </div>
  );
};

export default AIInsights;
