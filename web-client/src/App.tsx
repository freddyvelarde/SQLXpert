import { useEffect } from "react";
import useTheme from "./hooks/setTheme";
import Connection from "./views/Connection";

function App() {
  const { themeState, storeThemeIntoLocalStorage, colors, switchTheme } =
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

  return (
    <>
      <button onClick={() => console.log(colors)}>Get Colors</button>
      <button onClick={switchTheme}>Switch Theme</button>
      <Connection />
    </>
  );
}

export default App;
