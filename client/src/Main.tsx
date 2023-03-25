import { Flex } from "@chakra-ui/react";
import { useEffect } from "react";
import { useRoutes } from "react-router-dom";
import { useAuthStore } from "./DAL/stores/auth";
import { DataObjectState } from "./DAL/stores/types";
import { mainRoutes } from "./routes";

function Main() {
   const routeResult = useRoutes(mainRoutes);
  return (
    <Flex
      flexDirection="column"
      width="100wh"
      height="100vh"
      backgroundColor="gray.200"
      justifyContent="center"
      alignItems="center"
    >
      {routeResult}
    </Flex>
  );
}

export { Main };