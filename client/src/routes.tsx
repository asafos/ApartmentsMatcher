import { RouteObject } from 'react-router-dom'
import { AddApartmentContainer } from './pages/App/AddApartment/AddApartment.container'
import { AppContainer } from './pages/App/App.container'
import { Home } from './pages/App/Home/Home'
import { LoginContainer } from './pages/Login/Login.container'

const mainRoutes = [
  {
    path: '/login',
    element: <LoginContainer />,
  },
  {
    path: '*',
    element: <AppContainer />,
  },
]

const appRoutes: RouteObject[] = [
  {
    path: '/apartment/add',
    element: <AddApartmentContainer />,
  },
  {
    path: '/',
    element: <Home />,
  },
]

export { mainRoutes, appRoutes }
