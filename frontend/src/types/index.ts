export interface WeatherData {
  time: string;
  temperature: number;
  humidity: number;
  windSpeed: number;
}

export interface StatsCardProps {
  title: string;
  value: string;
  icon: React.ComponentType<{ className?: string }>;
  iconColor: string;
}

export interface InputProps {
  label: string;
  type?: string;
  value: string;
  onChange: (e: React.ChangeEvent<HTMLInputElement>) => void;
  placeholder?: string;
}

export interface HeaderProps {
  onLogout: () => void;
}

export interface WeatherChartProps {
  data: WeatherData[];
}

export interface LoginPageProps {
  onLogin: () => void;
  onNavigateToRegister: () => void;
}

export interface RegisterPageProps {
  onRegister: () => void;
  onNavigateToLogin: () => void;
}

export interface DashboardProps {
  onLogout: () => void;
}

export interface LandingPageProps {
  onGetStarted: () => void;
}

export interface Insight {
  type: string;
  message: string;
  color: string;
}

export interface Feature {
  icon: React.ComponentType<{ className?: string }>;
  title: string;
  description: string;
}

export type PageType = "landing" | "login" | "register" | "dashboard";
