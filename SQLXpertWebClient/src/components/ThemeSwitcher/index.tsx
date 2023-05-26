import moon from "../../assets/moon.png";
import sun from "../../assets/sun.png";
import useTheme from "../../hooks/setTheme";
import { ThemeSwitcherStyled } from "./theme.switcher.styles";

export default function ThemeSwitcher() {
  const { themeState, switchTheme } = useTheme();

  return (
    <ThemeSwitcherStyled>
      {themeState ? (
        <img onClick={switchTheme} src={moon} alt="moon" />
      ) : (
        <img onClick={switchTheme} src={sun} alt="sun" />
      )}
    </ThemeSwitcherStyled>
  );
}
