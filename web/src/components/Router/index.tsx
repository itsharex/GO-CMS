import { useNavigate, useRoutes } from 'react-router-dom';
import { useAppSelector } from '@/store';
import { builderMenuRoutes } from '@/utils';
import { useEffect } from 'react';
import routes from '@/router';
import NotFont from '@/pages/NotFont';
import { constants } from '@/constant';

export const Router = () => {
  const navigate = useNavigate();
  const { menus, token } = useAppSelector((state) => state.UserStore);
  routes[1].children = builderMenuRoutes(menus) as any;
  routes[1].children?.push({
    path: '*',
    element: <NotFont />,
  });
  useEffect(() => {
    if (!token) {
      navigate(constants.routePath.login, { replace: true });
    }
  }, []);
  return useRoutes(routes);
};