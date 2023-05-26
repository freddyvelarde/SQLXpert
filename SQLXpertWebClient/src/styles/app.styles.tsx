import styled from "styled-components";
import { ColorProps } from "../interfaces/colors";

export const AppStyles = styled.div<ColorProps>`
  background: ${({ colors }) => colors.background};
  color: ${({ colors }) => colors.text};
  transition: 0.3s ease;
  min-height: 100vh;
`;
