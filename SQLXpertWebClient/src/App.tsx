import { useEffect } from "react";
import useTheme from "./hooks/setTheme";
import { RouterProvider, createBrowserRouter } from "react-router-dom";
import ErrorPage from "./pages/ErrorPage";
import Connection from "./pages/Connection";
import Dashboard from "./pages/Dashboard";
import useDatabases from "./hooks/useDatabases";
import Main from "./pages/Main";
import { AppStyles } from "./styles/app.styles";
import ThemeSwitcher from "./components/ThemeSwitcher";

function App() {
  const router = createBrowserRouter([
    {
      path: "/",
      element: <Main />,
      errorElement: <ErrorPage />,
    },
    {
      path: "/connection",
      element: <Connection />,
    },
    {
      path: "/:database",
      element: <Dashboard />,
    },
  ]);

  const { themeState, storeThemeIntoLocalStorage, colorPalette, fonts } =
    useTheme();
  useEffect(() => {
    const data = localStorage.getItem("theme");
    if (data !== null) {
      storeThemeIntoLocalStorage(JSON.parse(data));
    }
    // eslint-disable-next-line
  }, []);

  useEffect(() => {
    localStorage.setItem("theme", JSON.stringify(themeState));
  }, [themeState]);

  // databases store
  const { setDatabasesIntoRedux, databases } = useDatabases();

  useEffect(() => {
    const data = localStorage.getItem("databases");
    if (data !== null) {
      setDatabasesIntoRedux(JSON.parse(data));
    }
    // eslint-disable-next-line
  }, []);

  useEffect(() => {
    localStorage.setItem("databases", JSON.stringify(databases));
  }, [databases]);

  return (
    <AppStyles colors={colorPalette} fonts={fonts}>
      <RouterProvider router={router} />
      <ThemeSwitcher />
    </AppStyles>
  );
}

export default App;
