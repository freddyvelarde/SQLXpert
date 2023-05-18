import { useEffect } from "react";
import useTheme from "./hooks/setTheme";
import Connection from "./views/Connection";
import { useDispatch } from "react-redux";
import { setTheme } from "./redux/features/themeSlices";

function App() {
  const theme = useTheme();
  const dispatch = useDispatch();

  // storing theme into local storage
  useEffect(() => {
    const data = localStorage.getItem("theme");
    if (data !== null) {
      dispatch(setTheme(JSON.parse(data)));
    }
    // eslint-disable-next-line
  }, []);

  useEffect(() => {
    localStorage.setItem("theme", JSON.stringify(theme.themeState));
  }, [theme]);

  return (
    <>
      <button onClick={() => console.log(theme.colors)}>Get Colors</button>
      <button onClick={theme.switchTheme}>Switch Theme</button>

      <Connection />
    </>
  );
}

export default App;
