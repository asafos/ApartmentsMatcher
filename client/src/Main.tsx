import { useRoutes } from "react-router-dom";
import { mainRoutes } from "./routes";

function Main() {
  const routeResult = useRoutes(mainRoutes);
  return (
    <main>
      {routeResult}
    </main>
  );
}

export { Main };