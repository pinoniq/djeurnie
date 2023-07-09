import {
    STATE,
} from './constants';
import { MouseEventHandler } from "react";

export type SidebarProps = {
    state: SidebarState,
    toggleState: MouseEventHandler,
};

export type SidebarState = STATE;