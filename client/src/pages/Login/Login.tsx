import {
    Heading,
    Button,
    Stack,
    chakra,
    Box,
    Avatar
} from "@chakra-ui/react";
import { FaGoogle, FaFacebook } from "react-icons/fa";

const CFaFacebook = chakra(FaFacebook);
const CFaGoogle = chakra(FaGoogle);

type Props = {
    onFacebookClick: () => void
    onGoogleClick: () => void
}

const Login = (props: Props) => {
    const { onFacebookClick, onGoogleClick} = props;
    return (
        <Stack
            flexDir="column"
            mb="2"
            justifyContent="center"
            alignItems="center"
        >
            <Avatar bg="teal.500" />
            <Heading color="teal.400">Welcome</Heading>
            <Box minW={{ base: "90%", md: "468px" }}>
                <Stack
                    flexDir="column"
                    spacing={4}
                    p="1rem"
                    backgroundColor="whiteAlpha.900"
                    boxShadow="md"
                >
                    <Button
                        borderRadius={0}
                        type="submit"
                        variant="solid"
                        colorScheme="blue"
                        width="full"
                        leftIcon={<CFaFacebook />}
                        onClick={onFacebookClick}
                    >
                        Login with Facebook
                    </Button>
                    <Button
                        borderRadius={0}
                        type="submit"
                        variant="solid"
                        colorScheme="red"
                        width="full"
                        leftIcon={<CFaGoogle />}
                        onClick={onGoogleClick}
                    >
                        Login with google
                    </Button>
                </Stack>
            </Box >
        </Stack >
    );
};

export { Login };
