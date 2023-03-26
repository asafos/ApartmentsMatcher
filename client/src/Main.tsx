import { Flex } from "@chakra-ui/react";
import { useEffect } from "react";
import { useRoutes } from "react-router-dom";
import { useAuthStore } from "./DAL/stores/auth";
import { DataObjectState } from "./DAL/stores/types";
import { mainRoutes } from "./routes";

function Main() {
  const fetchUser = useAuthStore((state) => state.fetchUser)
  const user = useAuthStore((state) => state.user)

  useEffect(() => {
    if (user.state === DataObjectState.NotStarted) {
      fetchUser()
    }
  }, [user.state])

  const routeResult = useRoutes(mainRoutes);
  return (
    <Flex
      width="100%"
      height="100vh"
      display="block"
      backgroundColor="gray.200"
    >
      {routeResult}
    </Flex>
  );
}

export { Main };