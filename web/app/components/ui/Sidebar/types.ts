import {
    STATE,
} from './constants';

export type SidebarProps = {
    state: SidebarState,
    toggleState: Function,
};

export type SidebarState = STATE;