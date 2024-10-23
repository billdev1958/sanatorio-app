import { createSignal, createContext, useContext, ParentComponent } from "solid-js";

// Define la interfaz del contexto de autenticación
interface AuthContextType {
  isAuthenticated: () => boolean;
  login: (token: string) => void;
  logout: () => void;
  token: () => string | null;
}

// Crea el contexto con la posibilidad de ser nulo
const AuthContext = createContext<AuthContextType>();

export const AuthProvider: ParentComponent = (props) => {
  const [token, setToken] = createSignal<string | null>(localStorage.getItem("authToken"));

  const login = (newToken: string) => {
    setToken(newToken);
    localStorage.setItem("authToken", newToken);
  };

  const logout = () => {
    setToken(null);
    localStorage.removeItem("authToken");
  };

  const isAuthenticated = () => {
    return !!token();
  };

  return (
    <AuthContext.Provider value={{ isAuthenticated, login, logout, token }}>
      {props.children}
    </AuthContext.Provider>
  );
};

// Hook para usar el contexto de autenticación, lanzando un error si no está definido
export const useAuth = () => {
  const context = useContext(AuthContext);
  if (!context) {
    throw new Error("useAuth debe ser usado dentro de un AuthProvider");
  }
  return context;
};
