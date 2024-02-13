import { useState, FormEvent, FC } from 'react';
import { HttpStatusCode } from 'axios';
import { useNavigate } from 'react-router-dom';
import { PasswordField } from '../components/PasswordField'
import {
  Box,
  Button,
  Container,
  FormControl,
  FormLabel,
  Flex,
  Heading,
  Input,
  Stack,
} from '@chakra-ui/react'
import { axiosClient } from '../providers/AxiosClientProvider';

export const Login: FC = () => {
  const [userName, setuserName] = useState("");
  const [password, setPassword] = useState("");
  const navigate = useNavigate();
  const [error, setError] = useState("");

  const handleSubmit = (event: FormEvent) => {
    event.preventDefault();
    
    axiosClient.post("/login", {
      userName: userName,
      password: password,
    })
    .then((response) => {
      navigate(`/${userName}/dashboard`, { replace: true })  //TODO
    })
    .catch((error) => {
      if (error.response?.status === HttpStatusCode.Unauthorized) {
        console.error(error);
        setError('ログイン情報が間違っています。');
      } else {
        console.error(error);
        setError('通信に失敗しました。');
      }
    });
  };

  return (
    <Flex w="100vw" h="100wh">
    <Container maxW="lg" py={{ base: '12', md: '24' }} px={{ base: '0', sm: '8' }} >
    <Stack spacing="8">
      <Stack spacing="6">
        <Stack spacing={{ base: '2', md: '3' }} textAlign="center">
          <Heading size={{ base: 'xs', md: 'sm' }}>Log in to your account</Heading>
        </Stack>
      </Stack>
      <Box
        py={{ base: '0', sm: '8' }}
        px={{ base: '4', sm: '10' }}
        bg={{ base: 'transparent', sm: 'bg.surface' }}
        boxShadow={{ base: 'none', sm: 'md' }}
        borderRadius={{ base: 'none', sm: 'xl' }}
      >
        <Stack spacing="6">
          <Stack spacing="5">
            <FormControl>
              <FormLabel htmlFor="email">User Name</FormLabel>
              <Input id="userName" type="email" value={userName} onChange={(e) => setuserName(e.target.value)}/>
            </FormControl>
            <PasswordField id="password" type="password" value={password} onChange={(e) => setPassword(e.target.value)}/>
          </Stack>
          {error && <div style={{ color: 'red' }}>{error}</div>}
          <Stack spacing="6">
            <Button onClick={handleSubmit}>Sign in</Button>
          </Stack>
        </Stack>
      </Box>
    </Stack>
  </Container>
  </Flex>
  );
};

export default Login;
