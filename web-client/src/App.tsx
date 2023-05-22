import { useEffect } from "react";
import useTheme from "./hooks/setTheme";
// import Connection from "./views/Connection";
import { RouterProvider, createBrowserRouter } from "react-router-dom";
import ErrorPage from "./pages/ErrorPage";
import Dinamic from "./pages/Dinamic";
import Connection from "./views/Connection";

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
      path: "/:id",
      element: <Dinamic />,
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
