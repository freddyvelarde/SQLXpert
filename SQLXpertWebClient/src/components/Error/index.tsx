import useTheme from "../../hooks/setTheme";
import { ErrorStyled } from "./error.styles";

interface ErrorProps {
  message: string;
}

export default function Error({ message }: ErrorProps) {
  const { colorPalette } = useTheme();

  return (
    <ErrorStyled colors={colorPalette}>
      <p>{message}</p>
    </ErrorStyled>
  );
}
