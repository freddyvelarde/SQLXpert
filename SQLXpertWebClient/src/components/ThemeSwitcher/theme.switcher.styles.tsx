import styled from "styled-components";

export const ThemeSwitcherStyled = styled.div`
  width: 25px;
  display: flex;
  justify-content: center;
  align-items: center;
  img {
    scale: 0.6;
    &:hover {
      scale: 0.7;
    }
    cursor: pointer;
    transition: 0.3s;
  }
`;
