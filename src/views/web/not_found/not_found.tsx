import { useRouteError } from "react-router-dom";

export default function Notfound() {
  const routerError: any = useRouteError();
  return <div className="not_found"></div>;
}
