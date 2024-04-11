package icu.aimerny.elix.registry;

import icu.aimerny.elix.Elix;
import icu.aimerny.elix.enchantment.DropEnchantment;
import net.minecraft.registry.Registries;
import net.minecraft.registry.Registry;
import net.minecraft.util.Identifier;

public class ModEnchantment {

    public static void init() {
        Registry.register(Registries.ENCHANTMENT, new Identifier(Elix.MOD_ID, "drop"), new DropEnchantment());
    }

}
