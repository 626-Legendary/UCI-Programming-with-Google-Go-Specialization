# ⭐ **Define Moore’s law and explain why it has now stopped being true**

*(Full-credit version with all required concepts highlighted)*

**Moore’s law** is the observation made by Gordon Moore in 1965 that the number of transistors on an integrated circuit doubles approximately every two years, leading to consistent improvements in computing performance and reductions in cost per transistor. For decades, this empirical trend held because engineers were able to continually shrink transistor dimensions through advances in lithography and materials engineering.

However, **Moore’s law has effectively stopped being true** because modern semiconductor manufacturing has reached multiple **fundamental physical limitations** that prevent continued exponential transistor scaling. These limits include:

---

## **1. Quantum tunneling (major factor — 4 points)**

As transistors shrink into the single-digit nanometer regime, the gate oxide becomes only a few atomic layers thick. At this scale, electrons can “tunnel” through the barrier due to quantum mechanics, causing uncontrollable leakage currents. This prevents further shrinking of the gate dielectric or channel because the device can no longer reliably turn off. Quantum tunneling therefore sets a hard lower bound on how small transistor features can be, directly breaking the core mechanism that enabled Moore’s law.

---

## **2. Heat dissipation and power density limits (3 points)**

Even though transistor counts grew, the power each chip consumed could not scale proportionally. Power density increased dramatically as features shrank, producing **excessive heat that modern cooling systems cannot dissipate**. This thermal wall is why processor clock speeds have remained around 3–5 GHz since the mid-2000s. Since power scales with capacitance, voltage, and frequency, and voltage cannot drop further, higher transistor counts no longer translate into usable performance—another fundamental violation of Moore’s law’s original assumptions.

---

## **3. Voltage scaling failure (1 point)**

Dennard scaling predicted that as transistors shrink, supply voltage (Vdd) can also shrink. But voltage cannot fall below the **threshold voltage** required for reliable switching and noise margin. Because voltage scaling stalled around 1 volt, shrinking transistors no longer reduces power proportionally. Without voltage scaling, further transistor miniaturization becomes energetically inefficient and impractical.

---

## **4. Lithography and manufacturing limits (1 point)**

Traditional optical lithography reached its physical limit as wavelengths of light could not resolve features smaller than the wavelength itself. Even with EUV lithography, patterning costs increased dramatically, masks became more complex, and yields decreased. These manufacturing constraints make further transistor shrinking economically unsustainable, contributing to the end of Moore’s law in practice.

---

## **5. Interconnect delays and resistance–capacitance (RC) limits (1 point)**

Although transistors became smaller and faster, the **wires connecting them did not scale as effectively**. As interconnects became thinner, resistance rose and signal propagation slowed due to RC delay. At advanced nodes, interconnect delay—not transistor switching—became the dominant bottleneck. This means that adding more transistors does not produce proportional performance gains, violating Moore’s law’s implied performance trajectory.

---

# ⭐ **Conclusion**

Moore’s law depended on the assumption that transistors could continue shrinking indefinitely, improving speed, efficiency, and cost. **Quantum tunneling**, **power-density/heat limits**, **voltage-scaling failure**, **lithography barriers**, and **interconnect delays** together form the set of physical constraints that now prevent further exponential scaling. As a result, Moore’s law is no longer valid as a predictor of future semiconductor progress, and the industry has shifted to new approaches such as chiplet architectures, 3D stacking, and domain-specific accelerators instead of relying on simple transistor miniaturization.

---