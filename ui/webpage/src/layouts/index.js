import styles from './index.css';
import { Link } from 'umi';

function BasicLayout(props) {
  console.log(props);
  let active = styles.menuItem + ' ' + styles.menuItemActive;
  let normal = styles.menuItem;
  let path = props.location.pathname;
  document.title = '⭐️序列战争⭐️';
  return (
    <div className={styles.normal}>
      <div className={styles.headerRow}>
        <Link to="/" className={props.location.pathname === '/' ? active : normal}>游玩</Link>
        <Link to="/config" className={(path === '/config')||(path === '/register')||(path === '/login') ? active : normal }>设置</Link>
        <Link to="/help" className={props.location.pathname === '/help' ? active : normal}>帮助</Link>
        <Link to="/about" className={props.location.pathname === '/about' ? active : normal}>关于</Link>
      </div>
      {props.children}
    </div>
  );
}

export default BasicLayout;
