import styled from "styled-components";
import { ColorProps } from "../../interfaces/colors";

export const ErrorStyled = styled.div<ColorProps>`
  width: auto;
  margin: 10px;
  display: flex;
  justify-content: center;
  aligh-items: center;
  background: ${({ colors }) => colors.foreground};
`;
