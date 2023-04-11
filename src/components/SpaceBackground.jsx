import React, { useState, useEffect, useRef } from "react";
import BG from "vanta/dist/vanta.fog.min";
import * as THREE from "three";

const SpaceBackground = () => {
    const [vantaEffect, setVantaEffect] = useState(0);
    const vantaRef = useRef(null);

    useEffect(() => {
        if (!vantaEffect) {
            setVantaEffect(
                BG({
                    el: vantaRef.current,
                    THREE: THREE,
                    mouseControls: true,
                    touchControls: true,
                    gyroControls: false,
                    minHeight: 1080.00,
                    minWidth: 200.00,
                    highlightColor: 0xffffff,
                    midtoneColor: 0x1863f2,
                    lowlightColor: 0x1800a4,
                    baseColor: 0xe3e3e3,
                    speed: 0.50,
                })
            );
        }
        return () => {
            if (vantaEffect) vantaEffect.destroy();
        };
    }, [vantaEffect]);
    return (
        <div ref={vantaRef}>
        </div>
    );
};

export default SpaceBackground;
