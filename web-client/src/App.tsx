import { useEffect } from "react";
import useTheme from "./hooks/setTheme";
import Connection from "./views/Connection";

function App() {
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
      <Connection />
    </>
  );
}

export default App;
