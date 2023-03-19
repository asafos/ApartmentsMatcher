import { Link, useRoutes } from "react-router-dom";
import { mainRoutes } from "./routes";

function App() {
  const routeResult = useRoutes(mainRoutes);
  return (
    <>
      <main>
        {routeResult}
      </main>
    </>
  );
}

export default App;