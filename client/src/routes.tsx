import { RouteObject } from 'react-router-dom'
import { AddApartmentContainer } from './pages/App/AddApartment/AddApartment.container'
import { AddApartmentPrefContainer } from './pages/App/AddApartmentPref/AddApartmentPref.container'
import { AppContainer } from './pages/App/App.container'
import { Home } from './pages/App/Home/Home'
import { ErrorPage } from './pages/Error/Error'
import { LoginContainer } from './pages/Login/Login.container'

const mainRoutes = [
  {
    path: '/login',
    element: <LoginContainer />,
  },
  {
    path: '/error',
    element: <ErrorPage />,
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
    path: '/apartmentPref/add',
    element: <AddApartmentPrefContainer />,
  },
  {
    path: '/',
    element: <Home />,
  },
]

export { mainRoutes, appRoutes }
