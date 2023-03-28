import { useRoutes } from 'react-router-dom'
import { User } from '../../DAL/services/auth/auth'
import { appRoutes } from '../../routes'

type Props = {
  user: User | null
  isLoading: boolean
}

function App(props: Props) {
  const { user, isLoading } = props
  const routeResult = useRoutes(appRoutes)

  if (isLoading) {
    return null
  }

  if (!user) {
    return null
  }

  return <>{routeResult}</>
}

export { App }
