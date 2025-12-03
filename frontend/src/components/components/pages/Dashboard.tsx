import React from "react";
import { TrendingUp, Droplets, Wind, Eye } from "lucide-react";
import Header from "../Layout/Header";
import StatsCard from "../Dashboard/StatsCard";
import WeatherChart from "../Dashboard/WeatherChart";
import AIInsights from "../Dashboard/AiInsights";
import { mockWeatherData } from "../../../data/mockData";
import type { DashboardProps, StatsCardProps } from "../../../types";

const Dashboard: React.FC<DashboardProps> = ({ onLogout }) => {
  const stats: StatsCardProps[] = [
    {
      title: "Temperatura",
      value: "25°C",
      icon: TrendingUp,
      iconColor: "text-orange-500",
    },
    {
      title: "Umidade",
      value: "50%",
      icon: Droplets,
      iconColor: "text-blue-500",
    },
    {
      title: "Vento",
      value: "18 km/h",
      icon: Wind,
      iconColor: "text-green-500",
    },
    {
      title: "Visibilidade",
      value: "10 km",
      icon: Eye,
      iconColor: "text-purple-500",
    },
  ];

  return (
    <div className="min-h-screen bg-gray-50">
      <Header onLogout={onLogout} />

      <main className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        {/* Stats Cards */}
        <div className="grid grid-cols-1 md:grid-cols-4 gap-6 mb-8">
          {stats.map((stat, index) => (
            <StatsCard
              key={index}
              title={stat.title}
              value={stat.value}
              icon={stat.icon}
              iconColor={stat.iconColor}
            />
          ))}
        </div>

        {/* Chart and AI Insights */}
        <div className="grid grid-cols-1 lg:grid-cols-3 gap-6">
          <WeatherChart data={mockWeatherData} />
          <AIInsights />
        </div>
      </main>
    </div>
  );
};

export default Dashboard;
