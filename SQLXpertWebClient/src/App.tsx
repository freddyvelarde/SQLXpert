import { useEffect } from "react";
import useTheme from "./hooks/setTheme";
// import Connection from "./views/Connection";
import { RouterProvider, createBrowserRouter } from "react-router-dom";
import ErrorPage from "./pages/ErrorPage";
import Connection from "./pages/Connection";
import Dashboard from "./pages/Dashboard";

function App() {
  const router = createBrowserRouter([
    {
      path: "/",
      element: <h1>hello world from the main route</h1>,
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

  const { themeState, storeThemeIntoLocalStorage } = useTheme();
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

  return (
    <>
      <RouterProvider router={router} />
    </>
  );
}

export default App;
