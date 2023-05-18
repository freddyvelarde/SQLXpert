import { useDispatch, useSelector } from "react-redux";
import { RootState } from "../redux/store";
import { setTheme } from "../redux/features/themeSlices";

import colors from "../styles/themes.json";

export default function useTheme() {
  const theme = useSelector((state: RootState) => state.theme.value);
  const dispatch = useDispatch();

  const switchTheme = () => {
    dispatch(setTheme(!theme));
  };

  return { colors: theme ? colors.light : colors.dark, switchTheme };
}
