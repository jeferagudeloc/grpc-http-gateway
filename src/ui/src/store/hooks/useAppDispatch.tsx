import { AnyAction, ThunkDispatch } from "@reduxjs/toolkit";
import { AppDispatch } from "../types/appDispatch";
import { useDispatch } from "react-redux";

type TypedDispatch<T> = ThunkDispatch<T, any, AnyAction>;

export const useAppDispatch = () => useDispatch<TypedDispatch<AppDispatch>>();
