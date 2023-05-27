import styled from "styled-components";
import { ColorProps, FontProps } from "../interfaces/colors";

export const AppStyles = styled.div<ColorProps & FontProps>`
  background: ${({ colors }) => colors.background};
  color: ${({ colors }) => colors.bodyText};
  font-family: ${({ fonts }) => fonts.text};
  transition: 0.3s ease;
  min-height: 100vh;
`;
