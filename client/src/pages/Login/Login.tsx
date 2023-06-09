import { Heading, Button, Stack, chakra, Box, Avatar, Flex } from '@chakra-ui/react'
import { useGoogleLogin } from '@react-oauth/google'
import { FaGoogle, FaFacebook } from 'react-icons/fa'
import FacebookLogin from '@greatsumini/react-facebook-login'

const CFaFacebook = chakra(FaFacebook)
const CFaGoogle = chakra(FaGoogle)

type Props = {
  onGoogleLogin: (authCode: string) => void
  onFacebookClick: (authCode: string) => void
}

const Login = (props: Props) => {
  const { onFacebookClick, onGoogleLogin } = props
  const login = useGoogleLogin({
    onSuccess: (authCode) => onGoogleLogin(authCode.code),
    flow: 'auth-code',
  })

  return (
    <Flex
      flexDirection='column'
      width='100%'
      height='100%'
      justifyContent='center'
      alignItems='center'
    >
      <Stack flexDir='column' mb='2' justifyContent='center' alignItems='center'>
        <Avatar bg='teal.500' />
        <Heading color='teal.400'>Welcome</Heading>
        <Box minW={{ base: '90%', md: '468px' }}>
          <Stack
            flexDir='column'
            spacing={4}
            p='1rem'
            backgroundColor='whiteAlpha.900'
            boxShadow='md'
          >
            <FacebookLogin
              appId='1156511058373400'
              onSuccess={(response) => {
                onFacebookClick(response.signedRequest)
              }}
              scope='email'
              render={({ onClick }) => (
                <Button
                  borderRadius={0}
                  type='submit'
                  variant='solid'
                  colorScheme='blue'
                  width='full'
                  leftIcon={<CFaFacebook />}
                  onClick={onClick}
                >
                  Login with Facebook
                </Button>
              )}
            />
            <Button
              borderRadius={0}
              type='submit'
              variant='solid'
              colorScheme='red'
              width='full'
              leftIcon={<CFaGoogle />}
              onClick={() => login()}
            >
              Login with google
            </Button>
          </Stack>
        </Box>
      </Stack>
    </Flex>
  )
}

export { Login }
